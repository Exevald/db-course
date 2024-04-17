import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {BranchList} from "./view/pages/BranchList/BranchList";


const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);

root.render(
    <React.StrictMode>
            <BrowserRouter>
                <Routes>
                    <Route path={"/"} element={<BranchList branches={[]}/>}></Route>
                </Routes>
            </BrowserRouter>
    </React.StrictMode>
);
