const express = require("express");
const bodyParser = require("body-parser");
const angularExpressions = require("angular-expressions");

const app = express();
const PORT = 8080;

app.use(bodyParser.json());

// Vulnerable endpoint
app.post("/parse", (req, res) => {
    const { expression } = req.body;

    try {
        const evalExpression = angularExpressions.compile(expression);
        const result = evalExpression();

        res.status(200).send({ result });
    } catch (error) {
        res.status(500).send({ error: error.message });
    }
});

// Start the server
app.listen(PORT, () => {
    console.log(`Vulnerable app listening at http://localhost:${PORT}`);
});
