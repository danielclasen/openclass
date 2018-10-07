import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SessionListComponent } from './session-list.component';
import {ParticipantListComponent} from "../participant-list/participant-list.component";
import {
  MatButtonModule,
  MatChipsModule,
  MatDividerModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule, MatListModule
} from "@angular/material";
import {MatExpansionModule} from "@angular/material/expansion";
import {AppComponent} from "../app.component";
import {GravatarModule} from "ngx-gravatar";
import {ReactiveFormsModule} from "@angular/forms";
import {HttpClientModule} from "@angular/common/http";
import {NoopAnimationsModule} from "@angular/platform-browser/animations";

describe('SessionListComponent', () => {
  let component: SessionListComponent;
  let fixture: ComponentFixture<SessionListComponent>;

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
    fixture = TestBed.createComponent(SessionListComponent);
    component = fixture.componentInstance;
    component.course = {id: 1, title: "Test", description: "Foo", lecturer: "Mr. Bar", price: 9000};
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
