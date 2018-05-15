import React, { Component } from "react";
import { Link } from "react-router-dom";

class Header extends Component {
  render() {
    return (
      <nav className="pa3 pa4-ns">
        <Link
          className="link dim black b f1 f-headline-ns tc db mb3 mb4-ns"
          to="/"
        >
          Lift
        </Link>
        <div className="tc pb3">
          <Link className="link dim gray f6 f5-ns dib mr3" to="/">
            Programs
          </Link>
          <Link className="link dim gray f6 f5-ns dib mr3" to="/">
            About
          </Link>
          <Link className="link dim gray f6 f5-ns dib mr3" to="/signup">
            Sign Up
          </Link>
          <Link className="link dim gray f6 f5-ns dib" to="/login">
            Log In
          </Link>
        </div>
      </nav>
    );
  }
}

export default Header;
