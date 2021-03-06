import React, { Component } from "react";
import api from "../api";

const tableStyle = "f6 w-100 mw8 center";
const headerStyle = "fw6 bb b--black-20 tl pb3 pr3 bg-white";
const cellStyle = "pv3 pr3 bb b--black-20";
const inputStyle = "pa2 input-reset ba bg-transparent br3 w-100";

const EditExercise = ({ exercise, onChange }) => {
  return (
    <tr>
      <td className={cellStyle}>
        <input
          type="number"
          name="sets"
          value={exercise.sets}
          onChange={onChange}
          className={inputStyle}
        />
      </td>
      <td className={cellStyle} value={exercise.reps}>
        <input
          type="number"
          name="reps"
          value={exercise.reps}
          onChange={onChange}
          className={inputStyle}
        />
      </td>
      <td className={cellStyle} value={exercise.weight}>
        <input
          type="number"
          name="weight"
          value={exercise.weight}
          onChange={onChange}
          className={inputStyle}
        />
      </td>
      <td className={cellStyle}>
        <input
          type="text"
          name="movement"
          value={exercise.movement}
          onChange={onChange}
          className={inputStyle}
        />
      </td>
    </tr>
  );
};

const EditWorkout = ({ workout, onChange, addExercise }) => {
  const { exercises, name } = workout;
  const rows = exercises.map((exercise, i) => (
    <EditExercise
      key={i}
      exercise={exercise}
      onChange={event => onChange(i, event)}
    />
  ));
  return (
    <React.Fragment>
      <h2>{name}</h2>
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
      <button
        className={"f6 link dim br3 ba ph3 pv2 mg2 dib black bg-lightgray"}
        onClick={addExercise}
      >
        Add Exercise
      </button>
    </React.Fragment>
  );
};

class NewProgram extends Component {
  constructor() {
    super();
    this.state = {
      program: {
        name: "",
        workouts: [
          {
            name: "Day 1",
            exercises: [
              { sets: "5", reps: "5", weight: "225", movement: "Bench" }
            ]
          }
        ]
      }
    };
  }

  handleChange = (workoutIndex, exerciseIndex, { target }) => {
    const { name, value } = target;
    var { program } = this.state;
    // TODO use spreads to make change of state more efficient
    program.workouts[workoutIndex].exercises[exerciseIndex][name] = value;
    this.setState({ program });
  };
  addExercise = workoutIndex => {
    var { program } = this.state;
    program.workouts[workoutIndex].exercises.push({});
    this.setState({ program });
  };

  addWorkout = () => {
    var { program } = this.state;
    program.workouts.push({
      name: `Day ${program.workouts.length + 1}`,
      exercises: []
    });
    this.setState({ program });
  };

  addProgram = async () => {
    const { program } = this.state;
    const payload = await api.newProgram(program);
    console.log(payload);
  };

  render() {
    const { program } = this.state;

    const exercises = program.workouts.map((workout, i) => {
      return (
        <EditWorkout
          key={i}
          workout={workout}
          onChange={(exerciseIndex, event) =>
            this.handleChange(i, exerciseIndex, event)
          }
          addExercise={() => this.addExercise(i)}
        />
      );
    });

    return (
      <div>
        {exercises}
        <button onClick={this.addWorkout}>Add Workout</button>
        <button onClick={this.addProgram}>Save</button>
      </div>
    );
  }
}

export default NewProgram;
