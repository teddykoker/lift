import React, { Component } from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";
import { logout } from "../actions";

const mapStateToProps = state => ({
  user: state.user
});

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(logout())
});

class Header extends Component {
  render() {
    const loggedIn = !!this.props.user;
    if (loggedIn) {
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
            <span
              onClick={() => this.props.logout()}
              className="link dim gray f6 f5-ns dib mr3"
            >
              Logout
            </span>
          </div>
        </nav>
      );
    }
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

export default connect(mapStateToProps, mapDispatchToProps)(Header);
