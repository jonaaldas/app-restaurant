import { useLocalSearchParams } from "expo-router";
import { View, Text } from "react-native";
import { Restaurant } from "@/types/restaurants";
import { useState } from "react";

export default function Restaurant() {
  const { id } = useLocalSearchParams(); 
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
  const restaurant = restaurants.find((res: Restaurant) => res.place_id === id);  
  return (
    <View>
      <Text>Restaurant12121 {id}</Text>
    </View>
  );
}