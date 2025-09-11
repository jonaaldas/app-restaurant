interface Location {
  lat: number;
  lng: number;
}

interface Photo {
  height: number;
  html_attributions: string[];
  photo_reference: string;
  width: number;
}

interface GoogleReviewsPhoto {
  height: number;
  html_attributions: string[];
  photo_reference: string;
  width: number;
}

interface GoogleReviewsReview {
  author_name: string;
  author_url: string;
  language: string;
  original_language: string;
  profile_photo_url: string;
  rating: number;
  relative_time_description: string;
  text: string;
  time: number;
  translated: boolean;
}

interface GoogleReviewsResult {
  photos: GoogleReviewsPhoto[];
  rating: number;
  reviews: GoogleReviewsReview[];
}

interface Restaurant {
  name: string;
  rating: number;
  photos: Photo[];
  location: Location;
  place_id: string;
  would_try: boolean;
  reviews: GoogleReviewsResult;
  formatted_address: string;
}

interface RestaurantId {
  place_id: string;
  would_try: boolean;
}

export type { Restaurant, RestaurantId };