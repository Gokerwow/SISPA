export interface TransactionItemRequest {
  customer_id: number;
  therapist_id: number;
  service_id: number;
}

export interface TransactionRequest {
  nota_number?: string;
  date: string;
  status: string;
  discount_amount: number;
  sub_total: number;
  grand_total: number;

  items: TransactionItemRequest[];
}

export interface Invoice {
  id: number;
  nota_number: string;
  sub_total: number;
  discount_amount: number;
  grand_total: number;
  status: 'pending' | 'paid' | 'cancelled';
  payment_method?: paymentMethod | null;
  date: string;
  created_at?: string | null;
}

export interface InvoiceItem {
  id: number;
  invoice_id: number;
  therapist_id: number;
  service_id: number;
  customer_id: number;
  price: number;
}

export interface InvoiceData extends Invoice {
  items: InvoiceItem[];
}

export interface UpdateStatusRequest {
  status: 'pending' | 'paid' | 'cancelled';
  payment_method?: paymentMethod | null;
}

export type paymentMethod = 'cash' | 'qris' | 'transfer' | 'other';
