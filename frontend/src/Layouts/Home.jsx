import { NavLink } from "react-router-dom";

const Home = () => {
    return (
        <div className="home">
            <h3>Helsinki city bike app</h3>
            <p>This is pre-assignment for Solita Dev Academy Finland (fall 2022).</p>

            <p>Instructions:<br/>
            <a href="https://github.com/solita/dev-academy-2022-fall-exercise">https://github.com/solita/dev-academy-2022-fall-exercise</a></p>

            <p><b>Journeys</b> - List all journals. <NavLink to={"/journeys"}>/journeys</NavLink></p>
            <p><b>Stations</b> - List all stations. <NavLink to={"/stations"}>/stations</NavLink></p>
        </div>
    )
}

export default Home;