import React, { Component } from "react";
import { Link } from "react-router-dom";

class Programs extends Component {
  render() {
    return (
      <div>
        <Link to="/newProgram">New Program</Link>
      </div>
    );
  }
}

export default Programs;
