import { useEffect } from "react";
import { useState } from "react";
import { ArrowRightSquareFill, Hash, Link45deg } from "react-bootstrap-icons";
import axios from "axios";
import { NavLink } from "react-router-dom";

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
                <h2>Stations</h2>
                <h4>List all stations</h4>
                <p>Open specific station with "Open" button.</p>

                <table className="table table-sm table-borderless align-middle text-center table-hover">
                    <thead className="table-dark align-middle">
                        <tr>
                            <th scope="col"><Hash /></th>
                            <th scope="col">Name</th>
                            <th scope="col">Address</th>
                            <th scope="col">Operator</th>
                            <th scope="col">Capacity</th>
                            <th scope="col"><Link45deg /></th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            stations.map((s) => (
                                <tr key={s.id}>
                                    <td>{s.id}</td>
                                    <td>{s.name}</td>
                                    <td className="text-left">
                                        {s.osoite},<br />
                                        {s.kaupunki != " " ? (<>{s.kaupunki}</>) : "Helsinki"
                                        }
                                    </td>
                                    <td>{s.operaattor}</td>
                                    <td>{s.kapasiteet}</td>
                                    <td><a class="btn btn-dark" href={`/stations/${s.id}`} role="button">Open</a></td>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </div >
        )
    }
};

export default Stations;