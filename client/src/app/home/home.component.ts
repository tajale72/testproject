import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  users:any;
  email:string;


  constructor(private http: HttpClient) { }
  public findUserByEmailId() {
    
  }

  ngOnInit(): void {
  }

}
