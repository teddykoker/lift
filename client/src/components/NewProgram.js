import React, { Component } from "react";

const tableStyle = "f6 w-100 mw8 center";
const headerStyle = "fw6 bb b--black-20 tl pb3 pr3 bg-white";
const cellStyle = "pv3 pr3 bb b--black-20";

class NewProgram extends Component {
  constructor() {
    super();
    this.state = {
      program: {
        name: "",
        exercises: [{ sets: "5", reps: "5", weight: "225", movement: "Bench" }]
      },
      columns: 10
    };
  }

  render() {
    const { program } = this.state;
    const rows = program.exercises.map((exercise, i) => (
      <tr key={i}>
        <td className={cellStyle}>
          <input type="number" value={exercise.sets} />
        </td>
        <td className={cellStyle} value={exercise.reps}>
          <input type="number" value={exercise.reps} />
        </td>
        <td className={cellStyle} value={exercise.weight}>
          <input type="number" value={exercise.weight} />
        </td>
        <td className={cellStyle}>
          <input type="text" value={exercise.movement} />
        </td>
      </tr>
    ));

    return (
      <div className="overflow-auto">
        <table className={tableStyle} cellSpacing="0">
          <thead>
            <tr>
              <th className={headerStyle}>Sets</th>
              <th className={headerStyle}>Reps</th>
              <th className={headerStyle}>Weight</th>
              <th className={headerStyle}>Movement</th>
            </tr>
          </thead>
          <tbody className="lh-copy">{rows}</tbody>
        </table>
      </div>
    );
  }
}

export default NewProgram;
