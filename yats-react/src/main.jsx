import './style.css';
import React, { useEffect, useState } from 'react';
import ReactDom from 'react-dom/client';
import User from './components/User.jsx'

class Main extends React.Component{
  constructor(props) {
    super(props); // Calls the parent constructor
    this.state = {
      count: 0 // Initialize local state
    };
  }

  render() {
    return (<div>
    <div className="main-container">
      <header>
          <div className="adminInfo">
            <h4 className="adminName">admin name</h4>
          </div>
      </header>
      <div className="nav-container">
        <nav className="nav">
          <div className="nav-upper-options">
            <div className="nav-option option-1">
              <h3>Dashboard</h3>
            </div>
            <div className="nav-option option-2">
              <h3>
                User List
              </h3>
            </div>
          </div>
        </nav>
      </div>
        <div className="report-container">
          <div className="report-header">
            <h2 className='centered'>User List</h2>
          </div>
          <div className="report-body">
            <tr className="report-topic-heading">
              <th className="t-op">First Name </th>
              <th className="t-op">Last Name  </th>
              <th className="t-op">Role       </th>
              <th className="t-op">Status     </th>
              <th className='t-op' colSpan="2"></th>
            </tr>
          
            <User />
          </div>
        </div>
    </div>
          <footer id="footer">v1</footer>
    </div>
    );
  }
}

export default Main
const container = document.getElementById("root")
const root = ReactDom.createRoot(container);
root.render(<Main />)