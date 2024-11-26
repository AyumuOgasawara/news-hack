import React from "react";
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { useState

 } from "react";
import 'bootstrap/dist/css/bootstrap.min.css';

export default () => {

  const [formData, setFormData] = useState({
    numArticles: 1,
    keyword: '',
    phoneNumber: ''
  });

  const handleChange = (e) => {
    const { id, value } = e.target;
    setFormData({ ...formData, [id]: id === 'numArticles' ? parseInt(value, 10) : value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevent default form submission behavior
    try {
      const response = await fetch('http://backend:5001/articles', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Error in submission');
      }

      const data = await response.json();
      console.log('Success:', data);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <>
    <h1 style={{fontWeight:"bold", fontFamily: "monospace"
    }}>news hack
    </h1>

    <Form onSubmit={handleSubmit}>
      <Form.Group>
    <Form.Label style={{fontFamily: "monospace"}}>Number of articles</Form.Label>

    <Form.Select
      aria-label="Default select example"
      onChange={handleChange}
      value={formData.numArticles}
      id="numArticles"
    >

    {Array.from({ length: 10 }, (_, i) => (
      <option key={i + 1} value = {i+1}>
            {i + 1}
          </option>
        ))}
    </Form.Select>
    </Form.Group>

    <Form.Group>
    <Form.Label style={{fontFamily: "monospace"}}>Keyword</Form.Label>
          <Form.Control
            aria-label="keyword"
            id="keyword"
            onChange={handleChange}
            value={formData.keyword}
          />
    </Form.Group>

    <Form.Group>
    <Form.Label style={{fontFamily: "monospace"}}>Phone number (with country code)</Form.Label>
    <Form.Control
            aria-label="Phone number"
            id="phoneNumber"
            onChange={handleChange}
            value={formData.phoneNumber}
          />    </Form.Group>

    <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>

  </>
);

}
