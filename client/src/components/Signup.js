import React, { Component } from "react";
import api from "../api";
import { signup } from "../actions";
import { connect } from "react-redux";

const mapStateToProps = state => ({ ...state });

const mapDispatchToProps = dispatch => ({
  signup: payload => dispatch(signup(payload))
});

class Signup extends Component {
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
    const payload = await api.signup(username, password);
    this.props.signup(payload);
    if (payload.token) {
      window.localStorage.setItem("jwt", payload.token);
      api.setToken(payload.token);
      this.props.history.push("/");
    }
  };

  render() {
    const { username, password } = this.state;
    return (
      <form
        className="measure center"
        onSubmit={this.onSubmit(username, password)}
      >
        <fieldset className="ba b--transparent ph0 mh0">
          <legend className="f4 fw6 ph0 mh0">Signup</legend>
          <div className="mt3">
            <label className="db fw6 lh-copy f6">Username</label>
            <input
              className="pa2 input-reset ba bg-transparent w-100"
              name="username"
              onChange={this.onChange}
            />
          </div>
          <div className="mv3">
            <label className="db fw6 lh-copy f6">Password</label>
            <input
              className="pa2 input-reset ba bg-transparent w-100"
              type="password"
              name="password"
              onChange={this.onChange}
            />
          </div>
        </fieldset>
        <div>
          <input
            className="b ph3 pv2 input-reset ba b--black bg-transparent grow pointer f6 dib"
            type="submit"
            value="Signup"
          />
        </div>
      </form>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Signup);
