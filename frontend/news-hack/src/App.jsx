import React from "react";
import Dropdown from 'react-bootstrap/Dropdown';
import { Form } from "react-bootstrap";
import DropdownButton from 'react-bootstrap/DropdownButton';
import 'bootstrap/dist/css/bootstrap.min.css';

export default () => {
  const [frequency, setFrequency] = React.useState("Choose frequency")
  const [time, setTime] = React.useState("12:00")

  const handleEveryday = () => {
    setFrequency("Everyday");
    console.log(frequency);
  };
  const handleEveryWeekday = () => {
    setFrequency("Every weekday");
    console.log(frequency);
  };
  const handleEveryWeekend = () => {
    setFrequency("Every weekend");
    console.log(frequency);
  };

  const handleTimeChange = (event) => {
    setTime(event.target.value);
    console.log(event.target.value);
  };


  return (
  <>
    <h1 style={{fontWeight:"bold", fontFamily: "monospace"
    }}>news hack</h1>

    <DropdownButton id="dropdown-basic-button" title={frequency} style={{fontWeight:"bold", fontFamily: "monospace", width: '100%'
    }}>
      <Dropdown.Item onClick={handleEveryday}>Everyday</Dropdown.Item>
      <Dropdown.Item onClick={handleEveryWeekday}>Every weekday</Dropdown.Item>
      <Dropdown.Item onClick={handleEveryWeekend}>Every weekend</Dropdown.Item>
    </DropdownButton>


    <div style={{ marginTop: '20px', width: '50%'}}>
        <Form.Group controlId="formTimeInput">
          <Form.Control
            type="time"
            value={time}
            onChange={handleTimeChange}
            className="basic-text"
          />
        </Form.Group>
      </div>
  </>


);
}
