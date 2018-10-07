import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ParticipantListComponent } from './participant-list.component';
import {
  MatButtonModule,
  MatChipsModule,
  MatDividerModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule, MatListModule
} from "@angular/material";
import {SessionListComponent} from "../session-list/session-list.component";
import {MatExpansionModule} from "@angular/material/expansion";
import {AppComponent} from "../app.component";
import {GravatarModule} from "ngx-gravatar";
import {ReactiveFormsModule} from "@angular/forms";
import {HttpClientModule} from "@angular/common/http";
import {NoopAnimationsModule} from "@angular/platform-browser/animations";

describe('ParticipantListComponent', () => {
  let component: ParticipantListComponent;
  let fixture: ComponentFixture<ParticipantListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        AppComponent,
        ParticipantListComponent,
        SessionListComponent
      ],
      imports: [
        NoopAnimationsModule,
        HttpClientModule,
        ReactiveFormsModule,
        GravatarModule,
        MatExpansionModule,
        MatChipsModule,
        MatListModule,
        MatDividerModule,
        MatIconModule,
        MatButtonModule,
        MatInputModule,
        MatFormFieldModule
      ]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ParticipantListComponent);
    component = fixture.componentInstance;
    component.session = {id: 1, courseId: 1, date: new Date(), location: "Virtual", maxStudents: 10};
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
