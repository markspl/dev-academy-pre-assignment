import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from "react-router-dom"
import { Route, Routes } from "react-router"

import "bootstrap/dist/css/bootstrap.min.css";
import './index.css';

import { Container } from "react-bootstrap"
import Header from './Layouts/Header';
import Footer from './Layouts/Footer';
import Home from './Layouts/Home';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <BrowserRouter>
    <React.StrictMode>
      <Container>

        <div className="App">
          <Header />
          
          <Routes>
            <Route>

              <Route path="/" element={<Home />} />

            </Route>
          </Routes>

          <Footer />
        </div>
        
      </Container>
    </React.StrictMode>
  </BrowserRouter>
);

