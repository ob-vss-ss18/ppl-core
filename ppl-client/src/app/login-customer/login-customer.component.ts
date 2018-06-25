import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login-costumer',
  templateUrl: './login-customer.component.html',
  styleUrls: ['./login-customer.component.scss']
})
export class LoginCustomerComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

  loginCustomer(e) {
    e.preventDefault();
    const username = e.target.elements[0].value;
    console.log(username);
  }
}
