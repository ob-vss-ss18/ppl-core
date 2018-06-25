import { Component, OnInit } from '@angular/core';
import { Apollo } from "apollo-angular";
import gql from "graphql-tag";
import { Observable } from "rxjs/Observable";
import "rxjs/add/operator/map";

class Ski {
  constructor(
    public id: string = "",
    public manufacturer: string = "",
    public modell: string = "",
    public length: number = 0,
    public bodyheight: number = 0,
    public bodyweight: number = 0,
    public price_new: number = 0,
    public availability: Boolean = true
  ) {}
}

@Component({
  selector: 'app-equipment',
  templateUrl: './equipment.component.html',
  styleUrls: ['./equipment.component.scss']
})
export class EquipmentComponent implements OnInit {

 
// It maintains list of Skis
skis: Ski[] = [];
// It maintains Ski Model
regModel: Ski;
// It maintains Ski form display status. By default it will be false.
showNew: Boolean = false;
// It will be either 'Save' or 'Update' based on operation.
submitType: string = 'Save';
// It maintains table row index based on selection.
selectedRow: number;

skiList: Array<any> = []; // List of skis
comments: Observable<any>;

constructor(private apollo: Apollo) {

}


  ngOnInit() {
    this.displaySkis();

  }

  // Get all Skis
  displaySkis() {
    const getSkis = gql`
      {
        Skis {
          id
          manufacturer
          modell
          length
          bodyheight
          bodyweight
          price_new
          availability
        }
      }
    `;

    this.apollo
      .watchQuery({
        query: getSkis,
        fetchPolicy: "network-only"
      })
      .valueChanges.map((result: any) => result.data.Skis)
      .subscribe(data => {
        this.skis = data;
      });
  }

  // This method associate to New Button.
  onNew() {
    // Initiate new ski.
    this.regModel = new Ski();
    // Change submitType to 'Save'.
    this.submitType = "Save";
    // display ski entry section.
    this.showNew = true;
  }

  // This method associate to Save Button.
  onSave() {
    if (this.submitType === "Save") {
      const saveSki = gql`
        mutation createSki(
          $manufacturer: String!
          $modell: String!
          $length: Int
          $bodyheight: Int
          $bodyweight: Int
          $price_new: Float!
          $availability: Boolean
        ) {
          createSki(
            manufacturer: $manufacturer
            modell: $modell
            length: $length
            bodyheight: $bodyheight
            bodyweight: $bodyweight
            price_new: $price_new
            availability: $availability
          ) {
            id
          }
        }
      `;
      this.apollo
        .mutate({
          mutation: saveSki,
          variables: {
            manufacturer: this.regModel.manufacturer,
            modell: this.regModel.modell,
            length: this.regModel.length,
            bodyheight: this.regModel.bodyheight,
            bodyweight: this.regModel.bodyweight,
            price_new: this.regModel.price_new,
            availability: this.regModel.availability
          }
        })
        .subscribe(
          ({ data }) => {
            this.displaySkis();
          },
          error => {
            console.log("there was an error sending the query", error);
          }
        );

      // Push Ski model object into Ski list.
      // this.skis.push(this.regModel);
    } else {
      const updateSki = gql`
        mutation updateSki(
          $id: ID!
          $manufacturer: String!
          $modell: String!
          $length: Int
          $bodyheight: Int
          $bodyweight: Int
          $price_new: Float!
          $availability: Boolean
        ) {
          updateSki(
            id: $id
            manufacturer: $manufacturer
            modell: $modell
            length: $length
            bodyheight: $bodyheight
            bodyweight: $bodyweight
            price_new: $price_new
            availability: $availability
          ) {
            id
          }
        }
      `;
      this.apollo
        .mutate({
          mutation: updateSki,
          variables: {
            id: this.selectedRow + 1,
            manufacturer: this.regModel.manufacturer,
            modell: this.regModel.modell,
            length: this.regModel.length,
            bodyheight: this.regModel.bodyheight,
            bodyweight: this.regModel.bodyweight,
            price_new: this.regModel.price_new,
            availability: this.regModel.availability
          }
        })
        .subscribe(
          ({ data }) => {
            console.log("got editdata", data);
            this.displaySkis();
          },
          error => {
            console.log("there was an error sending the query", error);
          }
        );
    }
    // Hide Ski entry section.
    this.showNew = false;
  }

  // This method associate to Edit Button.
  onEdit(index: number) {
    
    // Assign selected table row index.
    this.selectedRow = index;
    // Initiate new Ski.
    this.regModel = new Ski();
    // Retrieve selected Ski from list and assign to model.
    this.regModel = Object.assign({}, this.skis[this.selectedRow]);

    // Change submitType to Update.
    this.submitType = "Update";
    // Display Ski entry section.
    this.showNew = true;
  }

  // This method associate to Delete Button.
  onDelete(index: number) {
    const deleteSki = gql`
      mutation deleteSki($id: ID!) {
        deleteSki(id: $id) {
          id
        }
      }
    `;
    this.apollo
      .mutate({
        mutation: deleteSki,
        variables: {
          id: index + 1
        }
      })
      .subscribe(
        ({ data }) => {
          console.log("got editdata", data);
          this.displaySkis();
        },
        error => {
          console.log("there was an error sending the query", error);
        }
      );
  }

  // This method associate toCancel Button.
  onCancel() {
    // Hide Ski entry section.
    this.showNew = false;
  }
}