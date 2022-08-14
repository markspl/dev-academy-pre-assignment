import Sidebar from "../Components/Sidebar"

const Home = () => {
    return (
        <div className="home">
            <div className="container-fluid">
                <div className="row">
                    <Sidebar />
                    <main className="col pt-4">
                        <h1>Hi, this is home</h1>
                        <p>Funny text</p>
                    </main>
                </div>
            </div>
        </div>
    )
}

export default Home;