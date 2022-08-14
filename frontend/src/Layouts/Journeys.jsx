import { useEffect } from "react";
import { useState } from "react";
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

    if (error.error) {
        return (
            <p>{error.message}</p>
        );
    } else {
        return (
            <div className="journeys">
                <table class="table table-sm table-borderless align-middle text-center">
                    <thead class="table-dark">
                        <tr>
                            <th scope="col">#</th>
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
                            journeys.map((journey) => (
                                <tr>
                                    <td>{journey.id}</td>
                                    <td>{journey.departure}</td>
                                    <td>{journey.return}</td>
                                    <td>{journey.departureStationName}</td>
                                    <td>{journey.returnStationName}</td>
                                    <td>{journey.distance}</td>
                                    <td>{journey.duration}</td>
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