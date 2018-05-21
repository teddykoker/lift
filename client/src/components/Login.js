import React, { Component } from "react";
import api from "../api";
import { login } from "../actions";
import { connect } from "react-redux";

const inputStyle = "pa2 input-reset ba bg-transparent br3 w-100";
const legendStyle = "f4 fw6 ph0 mh0";
const labelStyle = "db fw4 lh-copy f6";

const mapStateToProps = state => ({ ...state });

const mapDispatchToProps = dispatch => ({
  login: payload => dispatch(login(payload))
});

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = { username: "", password: "" };
  }
  onChange = event => {
    const { target } = event;
    this.setState({
      [target.name]: target.value
    });
  };

  onSubmit = (username, password) => async event => {
    event.preventDefault();
    const payload = await api.login(username, password);
    this.props.login(payload);
    if (payload.token) {
      window.localStorage.setItem("jwt", payload.token);
      api.setToken(payload.token);
      this.props.history.push("/");
    }
  };

  render() {
    const { username, password } = this.state;
    const { error } = this.props;
    return (
      <form
        className="measure center"
        onSubmit={this.onSubmit(username, password)}
      >
        {error}
        <fieldset className="ba b--transparent ph0 mh0">
          <legend className={legendStyle}>Login</legend>
          <div className="mt3">
            <label className={labelStyle}>Username</label>
            <input
              className={inputStyle}
              type="text"
              name="username"
              onChange={this.onChange}
            />
          </div>
          <div className="mv3">
            <label className={labelStyle}>Password</label>
            <input
              className={inputStyle}
              type="password"
              name="password"
              onChange={this.onChange}
            />
          </div>
        </fieldset>
        <div>
          <input
            className="input-reset f6 link dim br3 ba ph3 pv2 mb2 dib black bg-lightgray"
            type="submit"
            value="Login"
          />
        </div>
        <div className="lh-copy mt3">
          <a href="#0" className="f6 link dim black db">
            Sign up
          </a>
        </div>
      </form>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Login);
