import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-demo',
  templateUrl: './demo.component.html',
  styleUrls: ['./demo.component.scss']
})
export class DemoComponent {
  stockResponse = '...';
  authResponse = '...';
  priceResponse = '...';
  leasingResponse = '...';
  billMailResponse = '...';
  customerResponse = '...';
  reservationResponse = '...';
  //für heroku notwendig
  //corsanywhere = 'https://cors-anywhere.herokuapp.com/';

  pplStock() {
    let self = this

    // cors-anywhere

    fetch('https://ppl-stock.herokuapp.com/')
      .then(function (response) {
        return response.text();
      }).then(function (text) {
        console.log(text);
        self.stockResponse = text;
      });
  }

  authStock() {
    let self = this
    fetch('https://ppl-auth.herokuapp.com/')
      .then(function (response) {
        return response.text();
      }).then(function (text) {
        console.log(text);
        self.authResponse = text;
      });
  }

  priceCalStock() {
    let self = this
    fetch('https://ppl-pricecalculator.herokuapp.com/')
      .then(function (response) {
        return response.text();
      }).then(function (text) {
        console.log(text);
        self.priceResponse = text;
      });
  }

  leasingStock() {
    let self = this
    fetch('https://ppl-leasing.herokuapp.com/')
      .then(function (response) {
        return response.text();
      }).then(function (text) {
        console.log(text);
        self.leasingResponse = text;
      });
  }

  billMailStock() {
    let self = this
    fetch('https://ppl-billingandmailing.herokuapp.com/')
      .then(function (response) {
        return response.text();
      }).then(function (text) {
        console.log(text);
        self.billMailResponse = text;
      });
  }

  customerStock() {
    let self = this
    fetch('https://ppl-customer.herokuapp.com/')
      .then(function (response) {
        return response.text();
      }).then(function (text) {
        console.log(text);
        self.customerResponse = text;
      });
  }

  reservationStock() {
    let self = this
    fetch('https://ppl-reservation.herokuapp.com/')
      .then(function (response) {
        return response.text();
      }).then(function (text) {
        console.log(text);
        self.reservationResponse = text;
      });
  }


}
