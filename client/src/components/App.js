import React from "react";
import { connect } from "react-redux";
import { Route, Switch, withRouter } from "react-router-dom";
import api from "../api";
import { onLoad } from "../actions";

import Header from "./Header";

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  onLoad: token => dispatch(onLoad(token))
});

const Home = (props) => <div>Home</div>;
const Login = (props) => <div>Login</div>;

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
      <div>
        <Header />
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/login" component={Login} />
        </Switch>
      </div>
    );
  }
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(App));
