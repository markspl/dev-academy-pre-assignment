import { useEffect } from "react";
import { useState } from "react";
import axios from "axios";
import { useParams } from "react-router";
import { Hash } from "react-bootstrap-icons";

const Station = () => {
    const [station, setStation] = useState([]);
    const [error, setError] = useState({});

    const { stationId } = useParams();

    useEffect(() => {
        axios({
            method: "get",
            url: `http://localhost:3000/api/stations/${stationId}`,
            headers: {
                "Content-Type": "application/json",
            }
        }).then(res => {
            setStation(res.data);
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
                <table className="table table-sm table-borderless align-middle text-center table-hover">
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
                                <tr key={station.id}>
                                <td>{station.id}</td>
                                    <td>{station.name}</td>
                                    <td className="text-left">
                                        {station.osoite},<br />
                                        {station.kaupunki != " " ? (<>{station.kaupunki}</>) : "Helsinki"
                                        }
                                    </td>
                                    <td>{station.operaattor}</td>
                                    <td>{station.kapasiteet}</td>
                                </tr>

                    </tbody>
                </table>
            </div>
        )
    }
};

export default Station;