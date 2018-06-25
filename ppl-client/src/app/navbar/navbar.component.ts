import { Component, OnInit } from '@angular/core';
import { GC_AUTH_TOKEN } from '../constants';
import { Router } from '@angular/router';


@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  isCollapsed: Boolean = true;

  constructor(private router: Router) { }

  ngOnInit() {
  }

  logout() {
    localStorage.removeItem(GC_AUTH_TOKEN);
    this.router.navigate(['/login']);
  }

}
