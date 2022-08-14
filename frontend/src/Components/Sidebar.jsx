const Sidebar = () => {
    return (
        <nav className="col-md-2 col-sm-12 pt-md-4 pt-sm-2 d-md-block sidebar shadow-sm">
            <div className="position-sticky sidebar-sticky">
                <ul className="nav flex-column">
                    <li class="nav-item">
                        <a className="nav-link text-dark" href="#"><span>Journeys</span></a>
                        <a className="nav-link text-dark" href="#">Stations</a>
                    </li>
                </ul>
            </div>
        </nav>
    )
}

export default Sidebar;