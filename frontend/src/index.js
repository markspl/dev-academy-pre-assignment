import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from "react-router-dom"
import { Route, Routes } from "react-router"
import { Row } from 'react-bootstrap';

import "bootstrap/dist/css/bootstrap.min.css";
import './index.css';

import Header from './Layouts/Header';
import Footer from './Layouts/Footer';
import Home from './Layouts/Home';
import Journeys from './Layouts/Journeys';
import Sidebar from './Components/Sidebar';
import Stations from './Layouts/Stations';

const app = ReactDOM.createRoot(document.getElementById('app'));
app.render(
  <BrowserRouter>
    <React.StrictMode>
      <Header />
      <div className="container-fluid">
        <Row>
          <Sidebar />
          <main className="col pt-4">
            <Routes>
              <Route>
                <Route path="/" element={<Home />} />
                <Route path="/journeys" element={<Journeys />} />
                <Route path="/stations" element={<Stations />} />
              </Route>
            </Routes>
          </main>
        </Row>
        <Footer />
      </div>
    </React.StrictMode>
  </BrowserRouter>
);

