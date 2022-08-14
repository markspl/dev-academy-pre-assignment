import { Bicycle } from "react-bootstrap-icons"

const Header = () => {
    return (
        <header className="d-flex flex-wrap justify-content-center border-bottom shadow-sm">
            <a href="/" className="align-items-center text-dark text-decoration-none">
            <span className="fs-4"><Bicycle /> City Bike App</span>
            </a>
        </header>
    )
}

export default Header;