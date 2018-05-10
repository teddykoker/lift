import React from "react";

class App extends React.Component {
  render() {
    return (
      <nav class="pa3 pa4-ns">
        <a
          class="link dim black b f1 f-headline-ns tc db mb3 mb4-ns"
          href="#"
          title="Home"
        >
          Lift
        </a>
        <div class="tc pb3">
          <a class="link dim gray f6 f5-ns dib mr3" href="#">
            Programs
          </a>
          <a class="link dim gray f6 f5-ns dib mr3" href="#" title="About">
            About
          </a>
          <a class="link dim gray f6 f5-ns dib mr3" href="#" title="Store">
            Sign Up
          </a>
          <a class="link dim gray f6 f5-ns dib" href="#" title="Contact">
            Log In
          </a>
        </div>
      </nav>
    );
  }
}

export default App;
