import {Component, Input, OnInit} from '@angular/core';
import {BehaviorSubject} from "rxjs/internal/BehaviorSubject";
import {Session} from "../model/session";
import {Person} from "../model/person";
import {SessionService} from "../service/session.service";
import {FormBuilder, Validators} from "@angular/forms";
import {log} from "util";

@Component({
  selector: 'app-participant-list',
  templateUrl: './participant-list.component.html',
  styleUrls: ['./participant-list.component.scss']
})
export class ParticipantListComponent implements OnInit {

  constructor(private sessionService: SessionService,
              private fb: FormBuilder) { }

  private _session = new BehaviorSubject<Session>(null);
  private persons: Person[] = [];

  private personForm = this.fb.group({
    'firstName': ['', Validators.required ],
    'lastName': ['', Validators.required ],
    'email': ['', Validators.required ]
  });


  // change data to use getter and setter
  @Input()
  set session(value) {
    // set the latest value for _data BehaviorSubject
    this._session.next(value);
  };

  get session() {
    // get the latest value from _data BehaviorSubject
    return this._session.getValue();
  }

  ngOnInit() {
    this._session
      .subscribe(() => {
        this.fetchParticipants()
      });

    // this.personForm = this.fb.group({} as Person)
  }

  private fetchParticipants() {
    this.sessionService.getParticipationsForSession(this.session.id)
      .subscribe((data: Person[]) => this.persons = [...data]);
  }

  // noinspection JSMethodCanBeStatic
  onPersonFormSubmit(person: Person) {
    this.sessionService.participateInSession(this.session.id, person).subscribe( () => {
      this.fetchParticipants()
    })
  }

}
