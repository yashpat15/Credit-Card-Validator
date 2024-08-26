function validateCard() {
    var cardNumber = document.getElementById('cardNumber').value;
    var resultDiv = document.getElementById('result');

    if (!cardNumber) {
        resultDiv.textContent = 'Please enter a credit card number.';
        resultDiv.style.color = 'red';
        return;
    }

    // Setup the request
    fetch('/validate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ number: cardNumber })
    })
    .then(response => response.json())
    .then(data => {
        if (data.valid) {
            resultDiv.textContent = 'Valid Credit Card Number';
            resultDiv.style.color = 'green';
        } else {
            resultDiv.textContent = 'Invalid Credit Card Number';
            resultDiv.style.color = 'red';
        }
    })
    .catch(error => {
        console.error('Error:', error);
        resultDiv.textContent = 'Error validating the card.';
        resultDiv.style.color = 'red';
    });
}
