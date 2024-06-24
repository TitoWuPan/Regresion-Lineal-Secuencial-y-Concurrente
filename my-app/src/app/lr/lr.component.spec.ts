import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LRComponent } from './lr.component';

describe('LRComponent', () => {
  let component: LRComponent;
  let fixture: ComponentFixture<LRComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [LRComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LRComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
