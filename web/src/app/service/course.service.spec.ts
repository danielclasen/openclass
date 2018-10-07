import { TestBed, inject } from '@angular/core/testing';

import { CourseService } from './course.service';
import {HttpClientModule} from "@angular/common/http";

describe('CourseService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CourseService],
      imports: [
        HttpClientModule,
      ]
    });
  });

  it('should be created', inject([CourseService], (service: CourseService) => {
    expect(service).toBeTruthy();
  }));
});
