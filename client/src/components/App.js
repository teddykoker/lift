import React from "react";
import { connect } from "react-redux";
import api, { setToken } from "../api";
import { onLoad } from "../actions";

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  onLoad: token => dispatch(onLoad(token))
});

class App extends React.Component {
  componentWillMount() {
    const token = window.localStorage.getItem("jwt");
    if (token) {
      api.setToken(token);
    }
    this.props.onLoad(token);
  }

  render() {
    return (
      <nav className="pa3 pa4-ns">
        <a
          className="link dim black b f1 f-headline-ns tc db mb3 mb4-ns"
          href="#"
          title="Home"
        >
          Lift
        </a>
        <div className="tc pb3">
          <a className="link dim gray f6 f5-ns dib mr3" href="#">
            Programs
          </a>
          <a className="link dim gray f6 f5-ns dib mr3" href="#" title="About">
            About
          </a>
          <a className="link dim gray f6 f5-ns dib mr3" href="#" title="Store">
            Sign Up
          </a>
          <a className="link dim gray f6 f5-ns dib" href="#" title="Contact">
            Log In
          </a>
        </div>
      </nav>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(App);
