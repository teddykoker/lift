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

const navLinkStyle = "link dim white dib mr3";

class Header extends Component {
  render() {
    const loggedIn = !!this.props.user;

    return (
      <header className="bg-black-90 fixed w-100 ph3 pv3 pv4-ns ph4-m ph5-l">
        <nav className="f6 fw4 ttu tracked">
          <Link className={navLinkStyle} to="/">
            Lift
          </Link>
          <Link className={navLinkStyle} to="/programs">
            Programs
          </Link>
          <Link className={navLinkStyle} to="/">
            About
          </Link>
          {loggedIn ? (
            <span onClick={() => this.props.logout()} className={navLinkStyle}>
              Logout
            </span>
          ) : (
            <React.Fragment>
              <Link className={navLinkStyle} to="/signup">
                Sign Up
              </Link>
              <Link className={navLinkStyle} to="/login">
                Log In
              </Link>
            </React.Fragment>
          )}
        </nav>
      </header>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Header);
