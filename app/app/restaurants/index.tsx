import { View, Text, ScrollView } from "react-native";
import RestaurantCard from "@/components/RestaurantCard";
import { Restaurant } from "@/types/restaurants";
import { useEffect, useState } from "react";

export default function Restaurants() {
  const [restaurants, setRestaurants] = useState<Restaurant[]>([{
    name: "Mendocino Farms",
    rating: 4.8,
    photos: [{
      height: 0,
      html_attributions: [],
      photo_reference: "https://unsplash.com/photos/empty-table-and-chairs-inside-building-ZgREXhl8ER0?utm_content=creditShareLink&utm_medium=referral&utm_source=unsplash",
      width: 0,
    }],
    location: { lat: 0, lng: 0 },
    place_id: "123",
    would_try: false,
    reviews: { photos: [{
      height: 0,
      html_attributions: [],
      photo_reference: "https://unsplash.com/photos/empty-table-and-chairs-inside-building-ZgREXhl8ER0?utm_content=creditShareLink&utm_medium=referral&utm_source=unsplash",
      width: 0,
    }], rating: 4.5, reviews: [] },
    formatted_address: "123 Main St, San Francisco, CA 94102, USA",
  },
  {
    name: "Mendocino Farms",
    rating: 4.8,
    photos: [{
      height: 0,
      html_attributions: [],
      photo_reference: "https://unsplash.com/photos/empty-table-and-chairs-inside-building-ZgREXhl8ER0?utm_content=creditShareLink&utm_medium=referral&utm_source=unsplash",
      width: 0,
    }],
    location: { lat: 0, lng: 0 },
    place_id: "1243",
    would_try: true,
    reviews: { photos: [{
      height: 0,
      html_attributions: [],
      photo_reference: "https://unsplash.com/photos/empty-table-and-chairs-inside-building-ZgREXhl8ER0?utm_content=creditShareLink&utm_medium=referral&utm_source=unsplash",
      width: 0,
    }], rating: 4.5, reviews: [] },
    formatted_address: "123 Main St, San Francisco, CA 94102, USA",
    }]); 
  return (
    <ScrollView style={{ flex: 1 }}>
      <Text style={{ fontSize: 24, fontWeight: 'bold', margin: 20 }}>Restaurants</Text>
      <Text style={{ margin: 20 }}>Found {restaurants.length} restaurants</Text>
      {restaurants.map((restaurant, index) => {
        console.log('Mapping restaurant', index, restaurant.name);
        return (
          <RestaurantCard key={restaurant.place_id} restaurant={restaurant} />
        );
      })}
    </ScrollView>
  );
}
