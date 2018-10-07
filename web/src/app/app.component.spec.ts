import { TestBed, async } from '@angular/core/testing';
import { AppComponent } from './app.component';
import {ParticipantListComponent} from "./participant-list/participant-list.component";
import {SessionListComponent} from "./session-list/session-list.component";
import {
  MatButtonModule, MatChipsModule,
  MatDividerModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule,
  MatListModule
} from "@angular/material";
import {MatExpansionModule} from "@angular/material/expansion";
import {GravatarModule} from "ngx-gravatar";
import {ReactiveFormsModule} from "@angular/forms";
import {HttpClientModule} from "@angular/common/http";
describe('AppComponent', () => {
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        AppComponent,
        ParticipantListComponent,
        SessionListComponent
      ],
      imports: [
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
  it('should create the app', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));
});
