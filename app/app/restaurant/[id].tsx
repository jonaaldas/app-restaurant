import React, { useEffect, useState, useRef } from "react";
import { Share } from "react-native";
import { useLocalSearchParams, router } from "expo-router";
import {
  View,
  Text,
  ScrollView,
  Image,
  TouchableOpacity,
  Dimensions,
  StyleSheet,
  SafeAreaView,
  NativeScrollEvent,
  NativeSyntheticEvent,
  Animated,
  Linking,
} from "react-native";
import { Restaurant } from "@/types/restaurants";
import { Ionicons } from "@expo/vector-icons";

const screenWidth = Dimensions.get("window").width;
const screenHeight = Dimensions.get("window").height;

export default function RestaurantDetail() {
  const { id } = useLocalSearchParams();
  const idParam = Array.isArray(id) ? id[0] : id;

  const scrollViewRef = useRef<ScrollView>(null);
  const mainScrollViewRef = useRef<ScrollView>(null);
  const [currentImageIndex, setCurrentImageIndex] = useState(0);
  const scrollY = useRef(new Animated.Value(0)).current;
  const headerHeight = screenHeight * 0.55;

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
              photoUri:
                "https://lh3.googleusercontent.com/a/default-user=s100-p-k-no-mo",
            },
          ],
          flagContentUri:
            "https://www.google.com/local/imagery/report/?cb_client=maps_api_places",
          googleMapsUri:
            "https://www.google.com/maps/place//data=!3m4!1e2!3m2!1sAF1QipNUcNnGRjpZdekgoMW5g3ns6F24qhy1as-7XN2I!2e10",
        },
        {
          name: "https://images.unsplash.com/photo-1555992336-03a23c7b20ee?w=800",
          widthPx: 800,
          heightPx: 600,
          authorAttributions: [
            {
              displayName: "Málà Project",
              uri: "https://maps.google.com/maps/contrib/123456789",
              photoUri:
                "https://lh3.googleusercontent.com/a/default-user=s100-p-k-no-mo",
            },
          ],
          flagContentUri:
            "https://www.google.com/local/imagery/report/?cb_client=maps_api_places",
          googleMapsUri:
            "https://www.google.com/maps/place//data=!3m4!1e2!3m2!1sAF1QipNUcNnGRjpZdekgoMW5g3ns6F24qhy1as-7XN2I!2e10",
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
        weekday_descriptions: [
          "Monday - Friday: 11:00 AM - 10:00 PM",
          "Saturday - Sunday: 10:00 AM - 11:00 PM",
        ],
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
              photoUri:
                "https://lh3.googleusercontent.com/a/default-user2=s100-p-k-no-mo",
            },
          ],
          flagContentUri:
            "https://www.google.com/local/imagery/report/?cb_client=maps_api_places",
          googleMapsUri:
            "https://www.google.com/maps/place//data=!3m4!1e2!3m2!1sAF1QipNUcNnGRjpZdekgoMW5g3ns6F24qhy1as-7XN2I!2e10",
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
        weekday_descriptions: [
          "Monday - Friday: 11:00 AM - 10:00 PM",
          "Saturday - Sunday: 10:00 AM - 11:00 PM",
        ],
        next_close_time: "11:00 PM",
      },
    },
  ]);

  const restaurant = restaurants.find(
    (res: Restaurant) => res.place_id === idParam
  );

  if (!restaurant) {
    return (
      <SafeAreaView style={styles.container}>
        <Text>Restaurant not found</Text>
      </SafeAreaView>
    );
  }

  const coverImages = restaurant.photos.map((photo) => photo.name);
  const [currentReviewIndex, setCurrentReviewIndex] = useState(0);
  const [reviewImageIndex, setReviewImageIndex] = useState(0);

  // Get current review or null if no reviews
  const currentReview =
    restaurant.reviews.reviews.length > 0
      ? restaurant.reviews.reviews[currentReviewIndex]
      : null;

  const onMomentumScrollEnd = (e: NativeSyntheticEvent<NativeScrollEvent>) => {
    const x = e.nativeEvent.contentOffset.x;
    const index = Math.round(x / screenWidth);
    setCurrentImageIndex(index);
  };

  const onReviewImageScrollEnd = (
    e: NativeSyntheticEvent<NativeScrollEvent>
  ) => {
    const x = e.nativeEvent.contentOffset.x;
    const index = Math.round(x / screenWidth);
    setReviewImageIndex(index);
  };

  const nextReview = () => {
    if (restaurant.reviews.reviews.length > 1) {
      setCurrentReviewIndex((prev) =>
        prev === restaurant.reviews.reviews.length - 1 ? 0 : prev + 1
      );
      setReviewImageIndex(0); // Reset image index when changing reviews
    }
  };

  const prevReview = () => {
    if (restaurant.reviews.reviews.length > 1) {
      setCurrentReviewIndex((prev) =>
        prev === 0 ? restaurant.reviews.reviews.length - 1 : prev - 1
      );
      setReviewImageIndex(0); // Reset image index when changing reviews
    }
  };

  const headerTranslateY = scrollY.interpolate({
    inputRange: [0, headerHeight],
    outputRange: [0, -headerHeight],
    extrapolate: "clamp",
  });

  const headerOpacity = scrollY.interpolate({
    inputRange: [0, headerHeight * 0.5, headerHeight],
    outputRange: [1, 0.5, 0],
    extrapolate: "clamp",
  });

  return (
    <View style={styles.container}>
      {/* Animated Header */}
      <Animated.View
        style={[
          styles.imageSliderContainer,
          {
            transform: [{ translateY: headerTranslateY }],
            opacity: headerOpacity,
          },
        ]}
      >
        <ScrollView
          ref={scrollViewRef}
          horizontal
          pagingEnabled
          showsHorizontalScrollIndicator={false}
          scrollEventThrottle={16}
          onMomentumScrollEnd={onMomentumScrollEnd}
        >
          {coverImages.map((imageUrl, index) => (
            <View key={index} style={styles.slideContainer}>
              <Image
                source={{ uri: imageUrl }}
                style={styles.image}
                resizeMode="cover"
              />
            </View>
          ))}
        </ScrollView>

        <TouchableOpacity
          style={styles.closeButton}
          onPress={() => router.back()}
        >
          <Ionicons name="close" size={24} color="white" />
        </TouchableOpacity>

        <TouchableOpacity style={styles.shareButton} >
          <Ionicons name="share-outline" size={24} color="white" />
        </TouchableOpacity>

        <View style={styles.pagination}>
          {coverImages.map((_, index) => (
            <View
              key={index}
              style={[
                styles.paginationDot,
                index === currentImageIndex && styles.paginationDotActive,
              ]}
            />
          ))}
        </View>
      </Animated.View>

      {/* Main Content with Scroll Listener */}
      <Animated.ScrollView
        ref={mainScrollViewRef}
        style={styles.contentContainer}
        showsVerticalScrollIndicator={false}
        onScroll={Animated.event(
          [{ nativeEvent: { contentOffset: { y: scrollY } } }],
          { useNativeDriver: true }
        )}
        scrollEventThrottle={16}
        contentContainerStyle={{ paddingTop: headerHeight }}
      >
        <View style={styles.restaurantInfo}>
          <View style={styles.logoContainer}>
            <Image
              source={{ uri: restaurant.photos[0]?.name }}
              style={styles.logo}
            />
          </View>

          <Text style={styles.restaurantName}>{restaurant.name}</Text>
          <Text style={styles.cuisineType}>{restaurant.formatted_address}</Text>
          <View style={styles.statusContainer}>
            <Text style={styles.closedText}>
              {restaurant.current_opening_hours.open_now ? "Open" : "Closed"}
            </Text>
            <Text style={styles.openTime}>
              • Opens {restaurant.current_opening_hours.weekday_descriptions[0]}
            </Text>
          </View>

          <View style={styles.actionButtons}>
            <TouchableOpacity
              style={styles.reserveButton}
              onPress={() => Linking.openURL(restaurant.google_maps_uri)}
            >
              <Ionicons name="map-outline" size={20} color="white" />
              <Text style={styles.googleLink}>Google Maps</Text>
            </TouchableOpacity>

            <TouchableOpacity
              style={styles.connectedButton}
              onPress={() => Linking.openURL(restaurant.website_uri)}
            >
              <Ionicons name="globe-outline" size={20} color="black" />
              <Text style={styles.website}>Website</Text>
            </TouchableOpacity>
          </View>

          {/* Reviews Section */}
          <View style={styles.reviewsSection}>
            <View style={styles.reviewsHeader}>
              <Text style={styles.reviewsTitle}>Reviews</Text>
              {restaurant.reviews.reviews.length > 1 && (
                <View style={styles.reviewNavigation}>
                  <TouchableOpacity
                    onPress={prevReview}
                    style={styles.navButton}
                  >
                    <Ionicons name="chevron-back" size={20} color="#666" />
                  </TouchableOpacity>
                  <Text style={styles.reviewCounter}>
                    {currentReviewIndex + 1} of{" "}
                    {restaurant.reviews.reviews.length}
                  </Text>
                  <TouchableOpacity
                    onPress={nextReview}
                    style={styles.navButton}
                  >
                    <Ionicons name="chevron-forward" size={20} color="#666" />
                  </TouchableOpacity>
                </View>
              )}
            </View>

            {currentReview ? (
              <View style={styles.reviewContainer}>
                {/* Review Images Carousel */}
                {restaurant.reviews.photos &&
                  restaurant.reviews.photos.length > 0 && (
                    <View style={styles.reviewImagesContainer}>
                      <ScrollView
                        horizontal
                        pagingEnabled
                        showsHorizontalScrollIndicator={false}
                        onMomentumScrollEnd={onReviewImageScrollEnd}
                        scrollEventThrottle={16}
                      >
                        {restaurant.reviews.photos.map((photo, index) => (
                          <View key={index} style={styles.reviewImageSlide}>
                            <Image
                              source={{ uri: photo.photo_reference }}
                              style={styles.reviewImage}
                              resizeMode="cover"
                            />
                          </View>
                        ))}
                      </ScrollView>

                      {restaurant.reviews.photos.length > 1 && (
                        <View style={styles.reviewImagePagination}>
                          {restaurant.reviews.photos.map((_, index) => (
                            <View
                              key={index}
                              style={[
                                styles.reviewImageDot,
                                index === reviewImageIndex &&
                                  styles.reviewImageDotActive,
                              ]}
                            />
                          ))}
                        </View>
                      )}
                    </View>
                  )}

                {/* Review Text */}
                <View style={styles.reviewTextContainer}>
                  <Text style={styles.reviewText}>{currentReview.text}</Text>
                  <View style={styles.reviewMeta}>
                    <Text style={styles.reviewAuthor}>
                      — {currentReview.author_name}
                    </Text>
                    <Text style={styles.reviewTime}>
                      {currentReview.relative_time_description}
                    </Text>
                  </View>
                </View>
              </View>
            ) : (
              <View style={styles.noReviewsContainer}>
                <Text style={styles.noReviewsText}>No reviews available</Text>
              </View>
            )}
          </View>
        </View>
      </Animated.ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "white",
  },
  imageSliderContainer: {
    height: screenHeight * 0.55,
    position: "absolute",
    top: 0,
    left: 0,
    right: 0,
    zIndex: 1,
  },
  slideContainer: {
    width: screenWidth,
    height: screenHeight * 0.55,
  },
  image: {
    width: "100%",
    height: "100%",
  },
  closeButton: {
    position: "absolute",
    top: 50,
    left: 20,
    backgroundColor: "rgba(0, 0, 0, 0.7)",
    borderRadius: 20,
    width: 40,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
  },
  shareButton: {
    position: "absolute",
    top: 50,
    right: 20,
    backgroundColor: "rgba(0, 0, 0, 0.7)",
    borderRadius: 20,
    width: 40,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
  },
  pagination: {
    position: "absolute",
    bottom: 20,
    alignSelf: "center",
    flexDirection: "row",
  },
  paginationDot: {
    width: 8,
    height: 8,
    borderRadius: 4,
    backgroundColor: "rgba(255, 255, 255, 0.5)",
    marginHorizontal: 4,
  },
  paginationDotActive: {
    backgroundColor: "white",
    width: 24,
  },
  contentContainer: {
    flex: 1,
    backgroundColor: "white",
  },
  restaurantInfo: {
    backgroundColor: "white",
    borderTopLeftRadius: 20,
    borderTopRightRadius: 20,
    padding: 20,
    marginTop: -20,
    minHeight: screenHeight,
  },
  logoContainer: {
    width: 80,
    height: 80,
    borderRadius: 16,
    overflow: "hidden",
    marginBottom: 16,
    marginTop: 16,
    backgroundColor: "#f0f0f0",
  },
  logo: {
    width: "100%",
    height: "100%",
  },
  restaurantName: {
    fontSize: 28,
    fontWeight: "bold",
    marginBottom: 8,
    color: "#000",
  },
  cuisineType: {
    fontSize: 16,
    color: "#666",
    marginBottom: 8,
  },
  statusContainer: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: 20,
  },
  closedText: {
    color: "#ff4444",
    fontSize: 16,
    fontWeight: "500",
  },
  openTime: {
    color: "#666",
    fontSize: 16,
    marginLeft: 8,
  },
  actionButtons: {
    flexDirection: "row",
    gap: 12,
    marginBottom: 24,
  },
  reserveButton: {
    flex: 1,
    backgroundColor: "#000",
    borderRadius: 25,
    paddingVertical: 14,
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
    gap: 8,
  },
  googleLink: {
    color: "white",
    fontSize: 16,
    fontWeight: "600",
  },
  connectedButton: {
    flex: 1,
    backgroundColor: "white",
    borderRadius: 25,
    paddingVertical: 14,
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
    gap: 8,
    borderWidth: 1.5,
    borderColor: "#000",
  },
  website: {
    color: "#000",
    fontSize: 16,
    fontWeight: "600",
  },
  reviewsSection: {
    marginTop: 24,
  },
  reviewsHeader: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    marginBottom: 16,
  },
  reviewsTitle: {
    fontSize: 20,
    fontWeight: "600",
    color: "#000",
  },
  reviewNavigation: {
    flexDirection: "row",
    alignItems: "center",
    gap: 12,
  },
  navButton: {
    width: 32,
    height: 32,
    borderRadius: 16,
    backgroundColor: "#f0f0f0",
    justifyContent: "center",
    alignItems: "center",
  },
  reviewCounter: {
    fontSize: 14,
    color: "#666",
  },
  reviewContainer: {
    backgroundColor: "#f8f8f8",
    borderRadius: 12,
    overflow: "hidden",
  },
  reviewImagesContainer: {
    position: "relative",
  },
  reviewImageSlide: {
    width: screenWidth - 72, // Account for card margins and padding
    height: 200,
  },
  reviewImage: {
    width: "100%",
    height: "100%",
  },
  reviewImagePagination: {
    position: "absolute",
    bottom: 12,
    alignSelf: "center",
    flexDirection: "row",
  },
  reviewImageDot: {
    width: 6,
    height: 6,
    borderRadius: 3,
    backgroundColor: "rgba(255, 255, 255, 0.5)",
    marginHorizontal: 3,
  },
  reviewImageDotActive: {
    backgroundColor: "white",
    width: 20,
  },
  reviewTextContainer: {
    padding: 16,
  },
  reviewText: {
    fontSize: 15,
    lineHeight: 22,
    color: "#333",
    marginBottom: 12,
  },
  reviewMeta: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
  },
  reviewAuthor: {
    fontSize: 14,
    fontWeight: "600",
    color: "#666",
  },
  reviewTime: {
    fontSize: 13,
    color: "#999",
  },
  noReviewsContainer: {
    backgroundColor: "#f8f8f8",
    borderRadius: 12,
    padding: 32,
    alignItems: "center",
  },
  noReviewsText: {
    fontSize: 16,
    color: "#666",
    fontStyle: "italic",
  },
});
