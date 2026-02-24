export interface FullCustomer {
  id: number
  name: string
  phone: string
  birth_date: string
  address: string
  gender: string
  note: string
  total_spend: number
  total_visit: number
  created_at: string
}

export interface Customer {
  name: string
  phone: string
  birth_date: string
  address: string
  gender: string
  note: string
}

