const express = require('express');
const fetch = require('node-fetch'); // You can use 'axios' if you prefer
const bodyParser = require('body-parser');

const app = express();
const PORT = 5000; // Change this port if needed

// Middleware for parsing JSON bodies
app.use(bodyParser.json());

// Route to receive form data
app.post('/your-backend-route', async (req, res) => {
  try {
    const { numArticles, keyword, phoneNumber } = req.body;

    // Send the data to another REST API on localhost
    const apiResponse = await fetch('http://localhost:8080/articles', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ numArticles, keyword, phoneNumber }),
    });

    if (!apiResponse.ok) {
      throw new Error(`Failed to forward request, status code: ${apiResponse.status}`);
    }

    const apiData = await apiResponse.json();
    res.status(200).json({ message: 'Data successfully sent to external API', data: apiData });

  } catch (error) {
    console.error('Error handling request:', error);
    res.status(500).json({ message: 'Server error', error: error.message });
  }
});

app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});
