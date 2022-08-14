import { useEffect } from "react";
import { useState } from "react";
import { CalendarDate, Clock, Hash } from "react-bootstrap-icons";
import axios from "axios";

const Journeys = () => {
    const [journeys, setJourneys] = useState([]);
    const [error, setError] = useState({})

    useEffect(() => {
        axios({
            method: "get",
            url: "http://localhost:3000/api/journeys",
            headers: {
                "Content-Type": "application/json",
            }
        }).then(res => {
            setJourneys(res.data);
        }).catch(err => {
            setError({ error: true, err, message: err.message });
        })
    }, [])

    // for console.log'ing
    useEffect(() => {
        if (error.error) {
            console.error(`[Journeys] Error: ${error.message}`);
            console.error(error.err);
        }
    }, [error])

    // Departure/Return format
    function formatDate(date) {
        const d = new Date(date);
        const h = ("0" + d.getHours()).slice(-2)   // Adds "0" at the start
        const m = ("0" + d.getMinutes()).slice(-2)
        const s = ("0" + d.getSeconds()).slice(-2)

        return (
            <>
                <CalendarDate /> {d.getDate()}.{d.getMonth()}.{d.getFullYear()}<br />
                <Clock /> {h}:{m}.{s}
            </>
        )
        //d.toLocaleDateString("fi-FI", options);
    }

    // Duration format
    function formatTime(sec) {
        const m = Math.floor(sec / 60);
        const s = sec - m * 60;
        return (
            <>
                {m}min {s ? (s + "s") : ""}
            </>
        )
    }

    if (error.error) {
        return (
            <p>{error.message}</p>
        );
    } else {
        return (
            <div className="journeys">
                <h2>Journeys</h2>
                <h4>List all journeys</h4>
                
                <table className="table table-sm table-borderless align-middle text-center">
                    <thead className="table-dark align-middle">
                        <tr>
                            <th scope="col"><Hash /></th>
                            <th scope="col">Departure Time</th>
                            <th scope="col">Return Time</th>
                            <th scope="col">Departure Station</th>
                            <th scope="col">Return Station</th>
                            <th scope="col">Distance</th>
                            <th scope="col">Duration</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            journeys.map((j) => (
                                <tr key={j.id}>
                                    <td>{j.id}</td>
                                    <td>{formatDate(j.departure)}</td>
                                    <td>{formatDate(j.return)}</td>
                                    <td>{j.departureStationName}</td>
                                    <td>{j.returnStationName}</td>
                                    <td>{(j.distance / 1000).toFixed(2)}km</td>
                                    <td>{formatTime(j.duration)}</td>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </div>
        )
    }
};

export default Journeys;