from rest_framework.decorators import api_view
from sqlalchemy.orm import relationship


@api_view(["GET", "POST"])
