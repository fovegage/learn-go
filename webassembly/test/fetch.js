fetch('https://www.revolve.com/r/ajax/GetProductData.jsp?product%5B%5D=LSPA-WD148')
    .then(response => response.json())
    .then(json => console.log(json))
    .catch(err => console.log('Request Failed', err));


fetch('https://www.rugsusa.com/rugsusa/rugs/rugs-usa-atrium-medallion-washable/Gray/200BIRV16A-P.html')
    .then(response => response.json())
    .then(json => console.log(json))
    .catch(err => console.log('Request Failed', err));
