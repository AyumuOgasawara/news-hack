import React from "react";
import Dropdown from 'react-bootstrap/Dropdown';
import DropdownButton from 'react-bootstrap/DropdownButton';
import 'bootstrap/dist/css/bootstrap.min.css';

export default () => {
  const [frequency, setFrequency] = React.useState("Choose frequency")

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
  return (
  <>
    <h1 style={{fontWeight:"bold", fontFamily: "monospace"
    }}>news hack</h1>

    <DropdownButton id="dropdown-basic-button" title={frequency} style={{fontWeight:"bold", fontFamily: "monospace"
    }}>
      <Dropdown.Item onClick={handleEveryday}>Everyday</Dropdown.Item>
      <Dropdown.Item onClick={handleEveryWeekday}>Every weekday</Dropdown.Item>
      <Dropdown.Item onClick={handleEveryWeekend}>Every weekend</Dropdown.Item>
    </DropdownButton>

  </>
);
}
