import { NgModule } from '@angular/core';
import { Routes, RouterModule, CanActivate } from '@angular/router';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { LoginCustomerComponent } from './login-customer/login-customer.component';
import { DemoComponent } from './demo/demo.component';
import { EquipmentComponent } from './equipment/equipment.component';
import { CustomerComponent } from './customer/customer.component';
import { RentalComponent } from './rental/rental.component';
import { HomeComponent } from './home/home.component';
import { AuthGuard } from './auth/auth-guard';



const routes: Routes = [
    {
        path: '',
        component: HomeComponent,
        canActivate: [AuthGuard] 
    },
    {
        path: 'login',
        component: LoginComponent
    },
    {
        path: 'login-customer',
        component: LoginCustomerComponent,
        canActivate: [AuthGuard] 
    },
    {
        path: 'equipment',
        component: EquipmentComponent,
        canActivate: [AuthGuard] 
    },
    {
        path: 'customer',
        component: CustomerComponent,
        canActivate: [AuthGuard] 
    },
    {
        path: 'rental',
        component: RentalComponent,
        canActivate: [AuthGuard] 
    },
    {
      path: '**',
      redirectTo: '',
    }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
