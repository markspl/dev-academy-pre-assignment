import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from "react-router-dom"
import { Route, Routes } from "react-router"

import "bootstrap/dist/css/bootstrap.min.css";
import './index.css';

import Header from './Layouts/Header';
import Footer from './Layouts/Footer';
import Home from './Layouts/Home';
import Sidebar from './Components/Sidebar';
import { Row } from 'react-bootstrap';

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
              </Route>
            </Routes>
          </main>
        </Row>
        <Footer />
      </div>
    </React.StrictMode>
  </BrowserRouter>
);

