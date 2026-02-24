export interface Therapist {
  id: number
  name: string
  phone: string
  address: string
  commission_rate: number
  joined_date: string | null
  exit_date?: string | null
  status: string
  created_at?: string
  updated_at?: string
}

export interface TherapistCreate {
  name: string
  phone: string
  address: string
  commission_rate: number
  joined_date: string | null
  exit_date?: string | null
  status: string
}
