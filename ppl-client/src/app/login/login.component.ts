import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth/auth.service';
import { Apollo } from "apollo-angular";
import { GC_USER_ID, GC_AUTH_TOKEN } from '../constants'

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(private router: Router,
    private authService: AuthService,
    private apollo: Apollo) {
}
  ngOnInit() {
  }

  loginUser(e) {
    e.preventDefault();
    var username = e.target.elements[0].value;
    var password = e.target.elements[1].value;
    console.log(username, password);
    localStorage.setItem(GC_AUTH_TOKEN, 'token');
    this.router.navigate(['/']);
  }
}
