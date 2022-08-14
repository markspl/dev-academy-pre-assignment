import { useEffect } from "react";
import { useState } from "react";
import { CalendarDate, Clock, ClockHistory, Hash } from "react-bootstrap-icons";
import axios from "axios";

const Stations = () => {
    const [stations, setStations] = useState([]);
    const [error, setError] = useState({})

    useEffect(() => {
        axios({
            method: "get",
            url: "http://localhost:3000/api/stations",
            headers: {
                "Content-Type": "application/json",
            }
        }).then(res => {
            setStations(res.data);
        }).catch(err => {
            setError({ error: true, err, message: err.message });
        })
    }, [])

    // for console.log'ing
    useEffect(() => {
        if (error.error) {
            console.error(`[Stations] Error: ${error.message}`);
            console.error(error.err);
        }
    }, [error])

    if (error.error) {
        return (
            <p>{error.message}</p>
        );
    } else {
        return (
            <div className="stations">
                <table className="table table-sm table-borderless align-middle text-center">
                    <thead className="table-dark align-middle">
                        <tr>
                            <th scope="col"><Hash /></th>
                            <th scope="col">Name</th>
                            <th scope="col">Address</th>
                            <th scope="col">Operator</th>
                            <th scope="col">Capacity</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            stations.map((s) => (
                                <tr key={s.id}>
                                    <td>{s.id}</td>
                                    <td>{s.name}</td>
                                    <td className="text-left">{s.osoite}<br/>{s.kaupunki}</td>
                                    <td>{s.operaattor}</td>
                                    <td>{s.kapasiteet}</td>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </div>
        )
    }
};

export default Stations;