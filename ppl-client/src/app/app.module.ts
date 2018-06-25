import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { NgxElectronModule } from 'ngx-electron';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { DemoComponent } from './demo/demo.component';
import { EquipmentComponent } from './equipment/equipment.component';
import { CustomerComponent } from './customer/customer.component';
import { RentalComponent } from './rental/rental.component';
import { LoginCustomerComponent } from './login-customer/login-customer.component';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { NavbarComponent } from './navbar/navbar.component';
import { FormsModule } from '@angular/forms';
import { HomeComponent } from './home/home.component';
import { GraphQLModule } from "./graphql.module";
import { AuthService } from './auth/auth.service';
import { AuthGuard } from './auth/auth-guard';



@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DemoComponent,
    EquipmentComponent,
    CustomerComponent,
    RentalComponent,
    LoginCustomerComponent,
    NavbarComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgxElectronModule,
    NgbModule.forRoot(),
    FormsModule,
    GraphQLModule
  ],
  providers: [AuthService, AuthGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
