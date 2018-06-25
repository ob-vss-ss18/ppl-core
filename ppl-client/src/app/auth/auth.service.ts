import {Injectable} from '@angular/core';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {GC_AUTH_TOKEN} from '../constants';

@Injectable()
export class AuthService {
  private userId: string = null;

  private _isAuthenticated = new BehaviorSubject(false);

  constructor() {
  }

    public isAuthenticated(): boolean {
    const token = localStorage.getItem(GC_AUTH_TOKEN);
    // TODO: Check whether the token is expired and return
    // true or false
    return !!token;
  }
}