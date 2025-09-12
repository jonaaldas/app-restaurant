import { View, Text, ScrollView, Pressable, TouchableOpacity, StyleSheet } from "react-native";
import RestaurantCard from "@/components/RestaurantCard";
import { Restaurant } from "@/types/restaurants";
import { useEffect, useState } from "react";
import { router } from "expo-router";
import { SafeAreaView } from "react-native-safe-area-context";
import { Ionicons } from "@expo/vector-icons";

export default function Restaurants() {
  const [restaurants, setRestaurants] = useState<Restaurant[]>([
    {
      name: "Málà Project 123123 12312",
      rating: 4.8,
      photos: [
        {
          name: "https://images.unsplash.com/photo-1555992336-03a23c7b20ee?w=800",
          widthPx: 800,
          heightPx: 600,
          authorAttributions: [
            {
              displayName: "Málà Project",
              uri: "https://maps.google.com/maps/contrib/123456789",
              photoUri: "https://lh3.googleusercontent.com/a/default-user=s100-p-k-no-mo"
            }
          ],
          flagContentUri: "https://www.google.com/local/imagery/report/?cb_client=maps_api_places",
          googleMapsUri: "https://www.google.com/maps/place//data=!3m4!1e2!3m2!1sAF1QipNUcNnGRjpZdekgoMW5g3ns6F24qhy1as-7XN2I!2e10"
        },
        {
          name: "https://images.unsplash.com/photo-1555992336-03a23c7b20ee?w=800",
          widthPx: 800,
          heightPx: 600,
          authorAttributions: [
            {
              displayName: "Málà Project",
              uri: "https://maps.google.com/maps/contrib/123456789",
              photoUri: "https://lh3.googleusercontent.com/a/default-user=s100-p-k-no-mo"
            }
          ],
          flagContentUri: "https://www.google.com/local/imagery/report/?cb_client=maps_api_places",
          googleMapsUri: "https://www.google.com/maps/place//data=!3m4!1e2!3m2!1sAF1QipNUcNnGRjpZdekgoMW5g3ns6F24qhy1as-7XN2I!2e10"
        },
      ],
      location: { lat: 0, lng: 0 },
      place_id: "123",
      would_try: false,
      reviews: {
        photos: [
          {
            height: 0,
            html_attributions: [],
            photo_reference:
              "https://images.unsplash.com/photo-1555992336-03a23c7b20ee?w=800",
            width: 0,
          },
          {
            height: 0,
            html_attributions: [],
            photo_reference:
              "https://images.unsplash.com/photo-1504674900247-0877df9cc836?w=800",
            width: 0,
          },
          {
            height: 0,
            html_attributions: [],
            photo_reference:
              "https://images.unsplash.com/photo-1540189549336-e6e99c3679fe?w=800",
            width: 0,
          },
        ],
        rating: 4.5,
        reviews: [
          {
            author_name: "John Doe",
            author_url: "",
            language: "en",
            original_language: "en",
            profile_photo_url: "",
            rating: 5,
            relative_time_description: "2 weeks ago",
            text: "Mala Project brings the fiery flavors of Sichuan dry pot to NYC, offering a customizable, spice-packed experience. With bold aromatics and a hip, cozy vibe, it's a must-visit for heat seekers.",
            time: Date.now(),
            translated: false,
          },
          {
            author_name: "Jane Smith",
            author_url: "",
            language: "en",
            original_language: "en",
            profile_photo_url: "",
            rating: 4,
            relative_time_description: "1 month ago",
            text: "Amazing authentic Sichuan cuisine! The customizable dry pot is fantastic and the spice level is perfect. Highly recommend for spice lovers.",
            time: Date.now(),
            translated: false,
          },
          {
            author_name: "Mike Johnson",
            author_url: "",
            language: "en",
            original_language: "en",
            profile_photo_url: "",
            rating: 5,
            relative_time_description: "3 weeks ago",
            text: "Best Sichuan food in East Village! The atmosphere is great and the food is incredibly flavorful. Don't miss their signature dishes!",
            time: Date.now(),
            translated: false,
          },
        ],
      },
      formatted_address: "Chinese / $$ / East Village",
      price_level: "$$",
      website_uri: "https://www.mala-project.com",
      google_maps_uri: "https://www.google.com/maps/place/Mala+Project",
      current_opening_hours: {
        open_now: true,
        weekday_descriptions: ["Monday - Friday: 11:00 AM - 10:00 PM", "Saturday - Sunday: 10:00 AM - 11:00 PM"],
        next_close_time: "11:00 PM",
      },
    },
    {
      name: "Mendocino Farms",
      rating: 4.8,
      photos: [
        {
          name: "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=800",
          widthPx: 800,
          heightPx: 600,
          authorAttributions: [
            {
              displayName: "Mendocino Farms",
              uri: "https://maps.google.com/maps/contrib/987654321",
              photoUri: "https://lh3.googleusercontent.com/a/default-user2=s100-p-k-no-mo"
            }
          ],
          flagContentUri: "https://www.google.com/local/imagery/report/?cb_client=maps_api_places",
          googleMapsUri: "https://www.google.com/maps/place//data=!3m4!1e2!3m2!1sAF1QipNUcNnGRjpZdekgoMW5g3ns6F24qhy1as-7XN2I!2e10"
        },
      ],
      location: { lat: 0, lng: 0 },
      place_id: "1243",
      would_try: true,
      reviews: {
        photos: [
          {
            height: 0,
            html_attributions: [],
            photo_reference:
              "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=800",
            width: 0,
          },
        ],
        rating: 4.5,
        reviews: [],
      },
      formatted_address: "American / $$ / West Village",
      price_level: "$$",
      website_uri: "https://www.mendocinofarms.com",
      google_maps_uri: "https://www.google.com/maps/place/Mendocino+Farms",
      current_opening_hours: {
        open_now: true,
        weekday_descriptions: ["Monday - Friday: 11:00 AM - 10:00 PM", "Saturday - Sunday: 10:00 AM - 11:00 PM"],
        next_close_time: "11:00 PM",
      },
    },
  ]);
  return (
    <SafeAreaView style={{ flex: 1 }}>
      <View style={{ flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center', margin: 20 }}>
      <TouchableOpacity
          style={styles.closeButton}
          onPress={() => router.back()}
        >
          <Ionicons name="close" size={24} color="white" />
        </TouchableOpacity>

        <TouchableOpacity style={styles.shareButton} >
          <Ionicons name="share-outline" size={24} color="white" />
        </TouchableOpacity> 
      </View>
      <Text style={{ fontSize: 24, fontWeight: 'bold', margin: 20 }}>Restaurants</Text>
      <Text style={{ margin: 20 }}>Found {restaurants.length} restaurants</Text>
      {restaurants.map((restaurant, index) => {
        console.log('Mapping restaurant', index, restaurant.name);
        return (
          <Pressable key={restaurant.place_id} onPress={() => {
            router.push({ pathname: "/restaurant/[id]", params: { id: restaurant.place_id } });
          }}>
            <RestaurantCard restaurant={restaurant} />
          </Pressable>
        );
      })}
    </SafeAreaView>
  );
}


const styles = StyleSheet.create({
  closeButton: {
    backgroundColor: "rgba(0, 0, 0, 0.7)",
    borderRadius: 20,
    width: 40,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
  },
  shareButton: {
    backgroundColor: "rgba(0, 0, 0, 0.7)",
    borderRadius: 20,
    width: 40,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
  },
});