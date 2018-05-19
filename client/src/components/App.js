import React from "react";
import { connect } from "react-redux";
import { Route, Switch, withRouter } from "react-router-dom";
import api from "../api";
import { onLoad } from "../actions";

import Login from "./Login";
import Header from "./Header";
import Signup from "./Signup";
import Programs from "./Programs";
import NewProgram from "./NewProgram";

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
      <React.Fragment>
        <Header />
        <div className="pt5 pt6-ns">
          <Switch>
            <Route exact path="/" component={Home} />
            <Route path="/login" component={Login} />
            <Route path="/signup" component={Signup} />
            <Route path="/programs" component={Programs} />
          </Switch>
        </div>
      </React.Fragment>
    );
  }
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(App));
