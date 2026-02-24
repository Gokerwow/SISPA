export interface Services {
  id: number
	name: string
	price_male: number
	price_female: number
	price_tourist: number
	is_active: boolean
	created_at: string
}

export interface ServicesCreate {
	name: string
	price_male: number
	price_female: number
	price_tourist: number
	is_active: boolean
}
