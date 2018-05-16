import React from "react";
import { connect } from "react-redux";
import { Route, Switch, withRouter } from "react-router-dom";
import api from "../api";
import { onLoad } from "../actions";

import Login from "./Login";
import Header from "./Header";
import Signup from "./Signup";

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  onLoad: (token, user) => dispatch(onLoad(token, user))
});

const Home = props => <div>Home</div>;

class App extends React.Component {
  async componentWillMount() {
    const token = window.localStorage.getItem("jwt");
    if (token) {
      api.setToken(token);
    }
    this.props.onLoad(token, token ? await api.currentUser() : null);
  }

  render() {
    return (
      <div>
        <Header />
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/signup" component={Signup} />
        </Switch>
      </div>
    );
  }
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(App));
