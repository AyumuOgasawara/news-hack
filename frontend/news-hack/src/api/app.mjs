import express from 'express';
import fetch from 'node-fetch';
import bodyParser from 'body-parser';
import cors from 'cors';

const app = express();
const PORT = 5001;

app.use(cors());
app.use(bodyParser.json());

app.post('/articles', async (req, res) => {
  try {
    const { numArticles, keyword, phoneNumber } = req.body;

    // Send the data to the GO API
    const apiResponse = await fetch('http://api:8080/articles', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "numArticles": numArticles, "keyword": keyword, "phoneNumber": phoneNumber }),
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
  console.log(`Server running on http://0.0.0.0:${PORT}`);
});
