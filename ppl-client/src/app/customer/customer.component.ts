import { Component, OnInit } from '@angular/core';
import { NgbDateStruct } from '@ng-bootstrap/ng-bootstrap';
import { Apollo } from "apollo-angular";
import gql from "graphql-tag";
import { Observable } from "rxjs/Observable";
import "rxjs/add/operator/map";

class Customer {
  constructor(
    public firstName: string = '',
    public lastName: string = '',
    public dob: NgbDateStruct = null,
    public email: string = '',
    // public password: string = '',
    public street: string = '',
    public housenumber: string = '',
    public postalcode: string = '',
    public city: string = '',
  ) {}
}

@Component({
  selector: 'app-customer',
  templateUrl: './customer.component.html',
  styleUrls: ['./customer.component.scss']
})
export class CustomerComponent implements OnInit {

// It maintains list of Customers
customers: Customer[] = [];
// It maintains customer Model
regModel: Customer;
// It maintains customer form display status. By default it will be false.
showNew: Boolean = false;
// It will be either 'Save' or 'Update' based on operation.
submitType: string = 'Save';
// It maintains table row index based on selection.
selectedRow: number;

customerList: Array<any> = []; // List of customers
comments: Observable<any>;

constructor(private apollo: Apollo) {

}


  ngOnInit() {
    this.displayCustomers();

  }

  // Get all customers
  displayCustomers() {
    const getCustomers = gql`
      {
        Customers {
          id
          firstName
          lastName
          dob
          email
          street
          housenumber
          postalcode
          city
        }
      }
    `;

    this.apollo
      .watchQuery({
        query: getCustomers,
        fetchPolicy: "network-only"
      })
      .valueChanges.map((result: any) => result.data.Customers)
      .subscribe(data => {
        this.customers = data;
      });
  }

  // This method associate to New Button.
  onNew() {
    // Initiate new customer.
    this.regModel = new Customer();
    // Change submitType to 'Save'.
    this.submitType = "Save";
    // display customer entry section.
    this.showNew = true;
  }

  // This method associate to Save Button.
  onSave() {
    var dateVal =
      this.regModel.dob.year.toString() +
      "-" +
      this.regModel.dob.month.toString() +
      "-" +
      this.regModel.dob.day.toString();
    if (this.submitType === "Save") {
      const saveCustomer = gql`
        mutation createCustomer(
          $firstName: String!
          $lastName: String!
          $dob: GQDate!
          $email: String!
          $street: String!
          $housenumber: String!
          $postalcode: String!
          $city: String!
        ) {
          createCustomer(
            firstName: $firstName
            lastName: $lastName
            dob: $dob
            email: $email
            street: $street
            housenumber: $housenumber
            postalcode: $postalcode
            city: $city
          ) {
            id
            dob
          }
        }
      `;
      this.apollo
        .mutate({
          mutation: saveCustomer,
          variables: {
            firstName: this.regModel.firstName,
            lastName: this.regModel.lastName,
            dob: new Date(dateVal),
            email: this.regModel.email,
            street: this.regModel.street,
            housenumber: this.regModel.housenumber,
            postalcode: this.regModel.postalcode,
            city: this.regModel.city
          }
        })
        .subscribe(
          ({ data }) => {
            this.displayCustomers();
          },
          error => {
            console.log("there was an error sending the query", error);
          }
        );

      // Push customer model object into customer list.
      // this.customers.push(this.regModel);
    } else {
      const updateCustomer = gql`
        mutation updateCustomer(
          $id: ID!
          $firstName: String!
          $lastName: String!
          $dob: GQDate!
          $email: String!
          $street: String!
          $housenumber: String!
          $postalcode: String!
          $city: String!
        ) {
          updateCustomer(
            id: $id
            firstName: $firstName
            lastName: $lastName
            dob: $dob
            email: $email
            street: $street
            housenumber: $housenumber
            postalcode: $postalcode
            city: $city
          ) {
            id
            city
          }
        }
      `;
      this.apollo
        .mutate({
          mutation: updateCustomer,
          variables: {
            id: this.selectedRow + 1,
            firstName: this.regModel.firstName,
            lastName: this.regModel.lastName,
            dob: new Date(dateVal),
            email: this.regModel.email,
            street: this.regModel.street,
            housenumber: this.regModel.housenumber,
            postalcode: this.regModel.postalcode,
            city: this.regModel.city,
          }
        })
        .subscribe(
          ({ data }) => {
            console.log("got editdata", data);
            this.displayCustomers();
          },
          error => {
            console.log("there was an error sending the query", error);
          }
        );
    }
    // Hide customer entry section.
    this.showNew = false;
  }

  // This method associate to Edit Button.
  onEdit(index: number) {
    
    // Assign selected table row index.
    this.selectedRow = index;
    // Initiate new customer.
    this.regModel = new Customer();
    // Retrieve selected customer from list and assign to model.
    this.regModel = Object.assign({}, this.customers[this.selectedRow]);
    const dob = new Date(this.customers[this.selectedRow].dob.toString());

    this.regModel.dob = {
      day: dob.getDate(),
      month: dob.getMonth() + 1,
      year: dob.getFullYear()
    };

    // Change submitType to Update.
    this.submitType = "Update";
    // Display customer entry section.
    this.showNew = true;
  }

  // This method associate to Delete Button.
  onDelete(index: number) {
    const deleteCustomer = gql`
      mutation deleteCustomer($id: ID!) {
        deleteCustomer(id: $id) {
          id
        }
      }
    `;
    this.apollo
      .mutate({
        mutation: deleteCustomer,
        variables: {
          id: index + 1
        }
      })
      .subscribe(
        ({ data }) => {
          console.log("got editdata", data);
          this.displayCustomers();
        },
        error => {
          console.log("there was an error sending the query", error);
        }
      );
  }

  // This method associate toCancel Button.
  onCancel() {
    // Hide customer entry section.
    this.showNew = false;
  }
}