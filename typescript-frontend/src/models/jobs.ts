export interface Job  {
    id: number;
    name: string;
    date_applied: string;
    job_type : string;
   
}


export interface JobResponse {
  success: boolean;
  message: string;
  data: Job[];
}

// The job Response is because the way i wrap my data inside my go backend